import json
import requests
import workdays
import datetime
import total_hour
from configutil import Configutil


def skip_if_weekend(func):
    def wrapper():
        weekend = [5, 6]  # 5,6 is the index for Saturday and Sunday for weekday()
        index_of_today = datetime.datetime.now().weekday()
        if index_of_today in weekend:
            return
        else:
            func()
    return wrapper


def get_remaining_hours_by_list_id(list_id):
    # read config.ini
    config = Configutil()
    trello_key = config.read_config('Trello', 'trello_key')
    trello_token = config.read_config('Trello', 'trello_token')

    url = 'https://api.trello.com/1/lists/%s/cards' % (list_id)
    params = dict(
        key=trello_key,
        token=trello_token
    )

    resp = requests.get(url, params)
    data = json.loads(resp.text)

    total_remaining = 0
    for i in data:
        card_id = i['id']
        total_remaining += get_remaining_by_card_id(card_id)
    return total_remaining


def get_remaining_key():
    # read config.ini
    config = Configutil()
    trello_key = config.read_config('Trello', 'trello_key')
    trello_token = config.read_config('Trello', 'trello_token')
    board_id = config.read_config('Trello', 'board_id')

    url = 'https://api.trello.com/1/boards/%s/pluginData' % (board_id)
    params = dict(
        key=trello_key,
        token=trello_token
    )

    resp = requests.get(url, params)
    data = json.loads(resp.text)
    '''
        [{
            "id": "57c6755ff8e2aab365ab49bb",
            "idPlugin": "56d5e249a98895a9797bebb9",
            "scope": "board",
            "idModel": "57a412f046b7a53cb7818a25",
            "value": "{\"fields\":[{\"n\":\"Estimation\",\"t\":1,\"b\":1,\"id\":22},{\"n\":\"Remaining\",\"t\":1,\"b\":1,\"id\":23}]}",
            "access": "shared"
        }]
    '''
    try:
        value_json = json.loads(data[0]['value'])
        for json_object in value_json['fields']:
            if is_remaining_string(json_object['n']):
                return json_object['id']
    except:
        return ""


def is_remaining_string(key_string):
    lower_case_string = key_string.lower()
    if lower_case_string.startswith('remain'):
        return True
    else:
        return False


def get_remaining_by_card_id(card_id):
    # read config.ini
    config = Configutil()
    trello_key = config.read_config('Trello', 'trello_key')
    trello_token = config.read_config('Trello', 'trello_token')

    url = 'https://api.trello.com/1/cards/%s/pluginData' % (card_id)
    params = dict(
        key=trello_key,
        token=trello_token
    )

    resp = requests.get(url, params)
    data = json.loads(resp.text)
    try:
        value_json = json.loads(data[0]['value'])
        remaining_key = get_remaining_key()
        key_string = str(remaining_key)
        return float(value_json['fields'][key_string])
    except:
        return 0


def get_unfinished_tasks_remaining_hours():
    # read config.ini
    config = Configutil()
    ready_to_do_list_id = config.read_config('Trello', 'ready_to_do_list_id')
    doing_list_id = config.read_config('Trello', 'doing_list_id')

    # calculate total remaining hours
    ready_to_do_list_remaining_hours = get_remaining_hours_by_list_id(ready_to_do_list_id)
    doing_list_remaining_hours = get_remaining_hours_by_list_id(doing_list_id)
    return str(ready_to_do_list_remaining_hours + doing_list_remaining_hours)


@skip_if_weekend
def write_remaining_hour_to_config():
    # read config.ini
    config = Configutil()
    # Get new remaining hour
    remaining = get_unfinished_tasks_remaining_hours()
    # Get origin remaining setting
    origin_remaining_config = config.read_config('Scrum', 'remaining')
    remainings_array = origin_remaining_config.split(',')
    # Get workdays
    workdays_array = workdays.get_sprint_workdays().split(',')
    # Get today information
    today_index = workdays_array.index(datetime.date.today().strftime("%Y-%m-%d"))
    workday_num = today_index + 1

    if len(remainings_array) is not workday_num:
        remainings_count = 0 if remainings_array[0] is '' else len(remainings_array)
        while remainings_count is not workday_num:
            # If loss some datas, fill its with current value
            if remainings_count is 0:
                config.write_config('Scrum', 'remaining', remaining)
                origin_remaining_config = remaining
            else:
                config.write_config('Scrum', 'remaining', origin_remaining_config + ',' + remaining)
                origin_remaining_config = origin_remaining_config + ',' + remaining
            remainings_count = remainings_count + 1
    else:
        remainings_array[today_index] = remaining
        new_remaining_list_string = ",".join(remainings_array)
        config.write_config('Scrum', 'remaining', new_remaining_list_string)


def write_unplanned_hour_to_config():

    new_hours_from_new_cards = total_hour.get_new_card_hour_and_set_config()
    config = Configutil()
    new_card_id = config.read_config('AllCard', 'new_card_id')
    new_card_id_list = new_card_id.split(',')
    unplanned_hours_list = get_unplanned_hour_list_from_config(config)
    unplanned_hours_list[-1] = new_hours_from_new_cards
    unplanned_hours_str = ','.join(unplanned_hours_list)

    config.write_config('Scrum', 'unplanned_hour', unplanned_hours_str)
    original_card_id_str = config.read_config('AllCard', 'card_id')
    card_id_str = ''
    if new_card_id:
        card_id_str = original_card_id_str + ',' + new_card_id
    else:
        card_id_str = original_card_id_str
    config.write_config('AllCard', 'card_id', card_id_str)
    config.write_config('AllCard', 'new_card_id', '')


def get_unplanned_hour_by_date(date):

    config = Configutil()
    workdays_list = config.read_config('Scrum', 'workdays').split(',')
    unplanned_hours_list = config.read_config('Scrum', 'unplanned_hour').split(',')
    try:
        index_of_specified_date = workdays_list.index(date)
        return unplanned_hours_list[index_of_specified_date]
    except IndexError:  # the exception(list out of index) indicates that there is missing values.
        return None
    except:
        return 0


def get_unplanned_hour_list_from_config(config):

    sprint_workdays_list = workdays.get_sprint_workdays().split(',')
    today = datetime.date.today().strftime("%Y-%m-%d")
    today_index = sprint_workdays_list.index(today)
    workdays_list = sprint_workdays_list[:today_index + 1]  # this list contains working days from starting day to today.

    unplanned_hours_str_from_config = config.read_config('Scrum', 'unplanned_hour')
    unplanned_hours_list = unplanned_hours_str_from_config.split(',')

    workday_and_unplanned_dict = {}  # a dictionary that has key value pairs with one workding and its unplanned hour.
    for workday in workdays_list:
        workday_and_unplanned_dict[workday] = get_unplanned_hour_by_date(workday)

    first_date_missing_unplanned_hour = ''
    for workday in workdays_list:
        if not workday_and_unplanned_dict[workday]:
            first_date_missing_unplanned_hour = workday
            break
    if first_date_missing_unplanned_hour != today:  # if day_start_to_miss is not today, unplanned hours are missed in config.
        for k, v in workday_and_unplanned_dict.items():
            if not v:
                workday_and_unplanned_dict[k] = 0
    else:
        workday_and_unplanned_dict[today] = 0
    new_unplanned_hours_list = []
    for workday in workdays_list:
        new_unplanned_hours_list.append(str(workday_and_unplanned_dict[workday]))  # convert the value into str or it cannot use string join method  below.
    return new_unplanned_hours_list