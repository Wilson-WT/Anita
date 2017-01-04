import ConfigParser

class Configutil:
    def __init__(self):
        self.config = 'config.ini'
        self.configparser = ConfigParser.ConfigParser()
        self.configparser.read(self.config)

    ''' read value from config '''
    def read_config(self, section, key):
        value = self.configparser.get(section, key)
        return value

    ''' write value to config '''
    def write_config(self, section, key, value):
        self.configparser.set(section, key, value)
        with open(self.config, 'wb') as configfile:
            self.configparser.write(configfile)
