
lang = "en";
langURL = getUrlParameter("lang");
if (langURL != undefined 
    && langURL != "" 
    && langURL != null
    && langURL.length > 1) {
    lang = langURL;
}
year = getUrlParameter("year");
month = getUrlParameter("month");
id = getUrlParameter("id");


var app = angular.module('app', ['BaseApp']);

function converToTweetURL(UserName, Id) {
    return "https://twitter.com/" + UserName + "/status/" + Id;
}


app.controller('PageController', function ($scope, $sce, $http, $rootScope, MetaService) {
    $scope.Nextpageid = "Data/" + lang + "/" + "tweets.json"
    $scope.TweetsList = []
    $scope.TweetsListIds = []
    $scope.endToTweet = 5;
    $scope.goGetMoreTweet = true;
    $scope.seeAlsoLogo = true;
    $scope.showMainTweet = true;
    $scope.showLoader = false;

    $scope.getTweetHTML = function (UserName, Id) {
        return $sce.trustAsHtml('<blockquote class="twitter-tweet"><a href="https://twitter.com/' + UserName + '/status/' + Id + '"></a></blockquote>');
    }

    $scope.twitterload = function () {
        setTimeout(function () {
            twttr.widgets.load();
        }, 500);
    };

    $scope.TweetsExist = function (tweetId) {
        for (var x = 0; x < $scope.TweetsListIds.length; x++) {
            if ($scope.TweetsListIds[x] == tweetId) {
                return true;
            }
        }
        return false;
    };

    $scope.GetMoreTweets = function () {
        $scope.goGetMoreTweet = false
        if ($scope.Nextpageid == "") {
            return;
        }
        $scope.showLoader = true;
        $http.get(AppBaseURL + $scope.Nextpageid, RequestConfig)
            .then(function onSuccess(response) {
                $scope.showLoader = false;
                res = response.data
                currantLength = $scope.TweetsList.length
                for (var i = 0; i < res.TweetsList.length; i++) {
                    if (!$scope.TweetsExist(res.TweetsList[i].Id)) {
                        $scope.TweetsList[currantLength + i] = $scope.getTweetHTML(res.TweetsList[i].UserName, res.TweetsList[i].Id);
                    }
                }
                $scope.twitterload();

                $scope.Nextpageid = res.Nextpageid;

                $scope.goGetMoreTweet = true;
            })
            .catch(function onError(response) {
                $scope.showLoader = false;
                console.log(response);
            });
    }

    $scope.GetMainTweet = function () {
        $scope.showLoader = true;
        $http.get(AppBaseURL + "Data/" + lang + "/" + year + "/" + month + "/" + id + ".json", RequestConfig)
            .then(function onSuccess(response) {
                $scope.showLoader = false;
                res = response.data
                $scope.tweet.Recommended = []
                $scope.tweet.HTML = $scope.getTweetHTML(res.UserName, res.Id);
                $scope.tweet.image = "Data/" + lang + "/" + year + "/" + month + "/" + id + ".png";
                $scope.TweetsListIds.push(res.Id);
                for (i = 0; i < res.Recommended.length; i++) {
                    $scope.TweetsListIds.push(res.Recommended[i].Id);
                    $scope.tweet.Recommended[i] = $scope.getTweetHTML(res.Recommended[i].UserName, res.Recommended[i].Id)
                }

                $scope.twitterload();
                $rootScope.metaservice = MetaService;
                $rootScope.metaservice.set($scope.tweet.image);

            })
            .catch(function onError(response) {
                $scope.showLoader = false;
                console.log(response);
            });
    }


    $scope.loadMore = function () {
        if ($scope.endToTweet <= $scope.TweetsList.length) {
            $scope.endToTweet += 5;
            $scope.twitterload();
        }
        if ($scope.endToTweet >= $scope.TweetsList.length && $scope.goGetMoreTweet) {
            $scope.GetMoreTweets();
        }
    };

    $scope.tweet = {}
    if (id != null && id != "") {
        $scope.showMainTweet = true;
        $scope.GetMainTweet();
    } else {
        $scope.showMainTweet = false;
        $scope.seeAlsoLogo = false;
    }

    $scope.GetMoreTweets();


});

