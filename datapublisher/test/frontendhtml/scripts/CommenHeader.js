


AppBaseURL = "http://localhost:9091/"
RequestConfig = {

}
function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
};


var BaseApp = angular.module('BaseApp', []);
BaseApp.directive('scrolled', [
    '$rootScope',
    '$window',
    function ($rootScope, $window) {
        return {
            link: function (scope, elem, attrs) {

                var scrollEnabled, loadData, scrollTrigger = .90, scrollEnabled = true;;
                $window = angular.element($window);
                if (attrs.lazyNoScroll != null) {
                    scope.$watch(attrs.lazyNoScroll, function (value) {
                        scrollEnabled = (value == true) ? false : true;
                    });
                }

                if ((attrs.lazyScrollTrigger != undefined) && (attrs.lazyScrollTrigger > 0 && attrs.lazyScrollTrigger < 100)) {
                    scrollTrigger = attrs.lazyScrollTrigger / 100;
                }

                loadData = function () {
                    var wintop = window.pageYOffset;
                    var docHeight = window.document.body.clientHeight;
                    var windowHeight = window.innerHeight//$window.height();
                    var triggered = (wintop / (docHeight - windowHeight));

                    if ((scrollEnabled) && (triggered >= scrollTrigger)) {
                        return scope.$apply(attrs.scrolled);
                    }
                };

                $window.on('scroll', loadData);
                scope.$on('$destroy', function () {
                    return $window.off('scroll', loadData);
                });
            }
        };
    }
]);

BaseApp.service('MetaService', function() { 
    var metaimage = '';
    return {
       set: function(newimage) { 
           metaimage = newimage;
       }, 
       metaimage: function() { return metaimage; }
    }
 });
