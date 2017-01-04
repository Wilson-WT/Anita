import os
import json
import requests
import ConfigParser
from configutil import Configutil


# Load Config
config = ConfigParser.ConfigParser()
config.read('config.ini')
READY_TO_DO_LIST_ID = config.get('Trello', 'ready_to_do_list_id')
DOING_LIST_ID = config.get('Trello', 'doing_list_id')
TRELLO_KEY = config.get('Trello', 'trello_key')
TRELLO_TOKEN = config.get('Trello', 'trello_token')


def get_total_hour():
    phases_id_list = [READY_TO_DO_LIST_ID, DOING_LIST_ID]
    total_hour = 0
    for list_id in phases_id_list:
        url = 'https://api.trello.com/1/lists/%s/cards' % (list_id)
        params = dict(
            key=TRELLO_KEY,
            token=TRELLO_TOKEN
        )
        resp = requests.get(url, params)
        datas = json.loads(resp.text)
        for data in datas:
            card_id = data['id']
            total_hour += get_card_estimate(card_id)
    return total_hour


def get_estimation_key():
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
            if is_estimation_string(json_object['n']):
                return json_object['id']
    except:
        return ""


def is_estimation_string(key_string):
    lower_case_string = key_string.lower()
    if lower_case_string.startswith('estimat'):
        return True
    else:
        return False


def get_card_estimate(card_id):
    ''' Get card hour '''
    url = 'https://api.trello.com/1/cards/%s/pluginData' % (card_id)
    params = dict(
        key=TRELLO_KEY,
        token=TRELLO_TOKEN
    )

    resp = requests.get(url, params)
    data = json.loads(resp.text)
    '''
    [{
      "id": "5805a179483d8c56dae27e26",
      "idPlugin": "56d5e249a98895a9797bebb9",
      "scope": "card",
      "idModel": "580581f3f8d90791f45b77bd",
      "value": "{\"fields\":{\"22\":\"5\",\"23\":\"4\"}}",
      "access": "shared"
    }]
    '''
    try:
        value_json = json.loads(data[0]['value'])
        estimation_key = get_estimation_key()
        key_string = str(estimation_key)
        return float(value_json['fields'][key_string])
    except:
        return 0


def get_new_card_hour_and_set_config():
    ''' Get newly added card hour and update config '''
    # get origin card id list
    config = ConfigParser.ConfigParser()
    config.read('config.ini')
    origin_card_id_list = config.get('AllCard', 'card_id').split(',')
    # get cards on board again
    phases_id_list = [READY_TO_DO_LIST_ID, DOING_LIST_ID]
    card_list = list()
    # get all phase's card_id
    for phase_id in phases_id_list:
        card_list_in_phase = get_list_cards_id(phase_id)
        card_list.extend(card_list_in_phase)
    # get difference
    new_card_list = list(set(card_list).difference(set(origin_card_id_list)))
    # Update card_id
    new_card_string = ",".join(new_card_list)
    # Caculate added hour
    added_hour = 0
    for new_card_id in new_card_list:
        added_hour += get_card_estimate(new_card_id)

    store_new_card_id(config, new_card_string)

    return str(added_hour)


def store_new_card_id(config, new_card_string):

    config.set('AllCard', 'new_card_id', new_card_string)
    with open('config.ini', 'wb') as configfile:
        config.write(configfile)


def store_all_card_id():
    phases_id_list = [READY_TO_DO_LIST_ID, DOING_LIST_ID]
    card_list = list()
    # get all phase's card_id
    for phase_id in phases_id_list:
        card_list_in_phase = get_list_cards_id(phase_id)
        card_list.extend(card_list_in_phase)
    # concat all card id : 1,2,3
    result = ','.join(card_list)
    # write to config.ini
    config = ConfigParser.ConfigParser()
    config.read('config.ini')
    config.set('AllCard', 'card_id', result)
    with open('config.ini', 'wb') as configfile:
        config.write(configfile)
    print 'store_all_card_id : ' + result


def get_list_cards_id(list_id):
    ''' Get a list's cards id '''
    url = 'https://api.trello.com/1/lists/%s/cards' % (list_id)
    # send request
    params = dict(
        key=TRELLO_KEY,
        token=TRELLO_TOKEN
    )
    resp = requests.get(url, params)
    datas = json.loads(resp.text)
    card_list = list()
    for data in datas:
        card_id = data['id']
        card_list.append(card_id)
    return card_list


def set_sprint_total_hour():
    ''' save sprint total hours to config.ini '''
    total_hour = get_total_hour()
    config = Configutil()
    config.write_config('Scrum', 'total_hour', total_hour)