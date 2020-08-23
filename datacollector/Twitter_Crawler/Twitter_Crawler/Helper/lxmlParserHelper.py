import os,sys


from lxml.html import fromstring as fromstringHTML , tostring as tostringHTML
import re

class lxmlParserHelper():

    @classmethod
    def stripList(self,MyList):
        try:
            result = []
            for element in MyList:
                stripElement = element.strip()
                if stripElement != "" :
                    result.append(stripElement)
            return result
        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            print ('stripList! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))


    @classmethod
    def lxmlParserFirst(self,element,pattarn):
        try:
            for x in fromstringHTML(element).xpath(pattarn):
                return x
            else :
                return None
        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            print ('lxmlParserFirst! {} - Line : {} - {} - {}'.format(fname,exc_tb.tb_lineno,str(e),exc_type))

    @classmethod
    def lxmlParserList(self,offer,pattarn):
        return fromstringHTML(offer).xpath(pattarn)