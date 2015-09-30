import webapp2
import json
import logging
import pickle
import urllib

from google.appengine.api import taskqueue, memcache, urlfetch
from google.appengine.api.taskqueue import Error
from google.appengine.api.urlfetch import Error as UrlfetchError



class Module1Handler(webapp2.RequestHandler):
  def get(self):
    logging.info("Get request to qnap-lab-002 module1.")
    self.response.write("Welcome to the qnap-lab-002 module1.")


class IncomeFileHandler(webapp2.RequestHandler):
  def get(self):
    logging.info("Get request to qnap-lab-002 module1/sayhi page.")
    self.response.write("sayhi.")

routes = [
  (r'/module1/sayhi', IncomeFileHandler),
  (r'/.*', Module1Handler)
]

router = webapp2.WSGIApplication(routes, debug=True)
