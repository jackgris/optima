'use strict';

/* Services */
angular.module('optima')
  .factory('Advertiser', AdvertiserService);

function AdvertiserService($resource){
    return $resource('static/js/json/:advertiserId.json', {}, {
        query: {method:'GET', params:{advertiserId:'advertisers'}, isArray:true}
    });
}
