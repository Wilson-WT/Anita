import ConfigParser, os, datetime
from configutil import Configutil

def get_sprint_workdays():
    # read config.ini
    config = Configutil()
    start_date_string = config.read_config('Scrum', 'start_date')
    end_date_string = config.read_config('Scrum', 'end_date')

    # convert string to datetime
    start_date_object = datetime.datetime.strptime(start_date_string, "%Y-%m-%d")
    end_date_object = datetime.datetime.strptime(end_date_string, "%Y-%m-%d")
    start_date = datetime.date(start_date_object.year, start_date_object.month, start_date_object.day)
    end_date = datetime.date(end_date_object.year, end_date_object.month, end_date_object.day)
    delta = end_date - start_date

    # build workdays string: 2016-11-01,2016-11-02,2016-11-03
    sprint_workday=""
    for i in range(delta.days + 1):
        date_test = start_date + datetime.timedelta(days=i)
        if date_test.isoweekday() != 6 and date_test.isoweekday() != 7:
            sprint_workday+=str(date_test) + ","

    return sprint_workday[:-1]

def write_workdays_to_config():
    config = Configutil()
    workdays = get_sprint_workdays()
    config.write_config('Scrum', 'workdays', workdays)
