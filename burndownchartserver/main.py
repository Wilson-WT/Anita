import schedule
import time
import remaining
import total_hour
import datetime
import workdays
from configutil import Configutil


def main():

    schedule.every().hour.do(remaining.write_remaining_hour_to_config)
    schedule.every().monday.at("16:30").do(remaining.write_unplanned_hour_to_config)
    schedule.every().tuesday.at("16:30").do(remaining.write_unplanned_hour_to_config)
    schedule.every().wednesday.at("16:30").do(remaining.write_unplanned_hour_to_config)
    schedule.every().thursday.at("16:30").do(remaining.write_unplanned_hour_to_config)
    schedule.every().friday.at("16:30").do(remaining.write_unplanned_hour_to_config)

    while True:
        schedule.run_pending()
        time.sleep(60)

if __name__ == "__main__":
    config = Configutil()
    # Process Total Hour
    start_date = config.read_config('Scrum', 'start_date')
    if start_date == datetime.date.today().strftime("%Y-%m-%d"):
        total_hour.set_sprint_total_hour()

    # Process Workdays
    workdays.write_workdays_to_config()
    total_hour.store_all_card_id()
    main()
