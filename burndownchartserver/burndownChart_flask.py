import ConfigParser
import os
import datetime
from flask import send_from_directory
from flask import Flask
from flask import Markup
from flask import Flask
from flask import render_template

app = Flask(__name__)


def check_config(config='config.ini'):

    config_parser = ConfigParser.ConfigParser()

    if os.path.isfile(config):
        config_parser.read(config)
    else:
        err_msg = '*** config.ini not exists...'
        print err_msg
        return (False, err_msg)

    if not config_parser.has_section('Scrum'):
        err_msg = '*** No [Scrum] in config.ini...'
        print err_msg
        return (False, err_msg)

    if not config_parser.has_option('Scrum', 'total_hour') or not config_parser.has_option('Scrum', 'remaining')\
       or not config_parser.has_option('Scrum', 'workdays'):
        err_msg = '*** Invalid option under [Scrum] in config.ini...'
        print err_msg
        return (False, err_msg)

    return (True, config_parser)


def get_unplanned_list_from_config(config='config.ini'):

    config_parser = ConfigParser.ConfigParser()
    config_parser.read(config)
    workdays = config_parser.get('Scrum', 'workdays').split(',')
    today_index = workdays.index(datetime.date.today().strftime("%Y-%m-%d"))
    workday_count = today_index + 1
    unplanned_hour_list = []
    unplanned_hour_list_from_config = config_parser.get('Scrum', 'unplanned_hour').split(',')

    for item in unplanned_hour_list_from_config:
        if item:
            unplanned_hour_list.append(float(item))

    if workday_count != len(unplanned_hour_list):
        # if the count of workdays is not equal to the len of unplanned hours, indicating something wrong with unplanned hours.
        unplanned_hour_list = []

    return unplanned_hour_list


@app.route("/")
def chart():

    config_check = check_config('config.ini')

    if not config_check[0]:
        err_msg = config_check[1] + '\n' + ' Check the config first... ***'
        return err_msg
    else:
        config_parser = config_check[1]

    total_hours = config_parser.get('Scrum', 'total_hour')
    workdays_string = config_parser.get('Scrum', 'workdays')
    workdays = workdays_string.split(',')
    ideal_line = []
    for _ in range(0, len(workdays)):
        ideal_line.append('null')
    ideal_line[0] = total_hours
    ideal_line[-1] = 0

    remaining_string = config_parser.get('Scrum', 'remaining')
    remaining = remaining_string.split(',')
    unplanned_list = get_unplanned_list_from_config()
    return render_template('chart.html', ideal_line=ideal_line, workdays=workdays, remaining=remaining, unplanned_list=unplanned_list)


@app.route('/favicon.ico')
def favicon():
    return send_from_directory(os.path.join(app.root_path, 'static'),
                               'favicon.ico', mimetype='image/vnd.microsoft.icon')


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5001)
