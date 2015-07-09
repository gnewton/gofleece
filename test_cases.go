package main

const smallestValidJsonMap = `{}`
const smallestValidJsonArray = `[]`

const validJsonArray0 = `
    [
   {
      "a":1,
      "b":3
   },
   {
      "a":1,
      "b":3
   }
]
`

const validJsonMap0 = `
    {
   "foo":{
      "a":1,
      "b":2
   },
   "bar":{
      "a":1,
      "b":2
   }
}`

const validJsonMap1 = `
    {
   "foo":{
      "a":1,
      "b":2
   }
}`

// from http://www.sitepoint.com/google-maps-json-file/
// Removed multiple point, like "point":new GLatLng(40.266044,-74.718479),
// Removed trailing comma after last entry
const google_maps_json_file = `{"markers": [
        {
            "point":"40.266044,-74.718479", 
            "homeTeam":"Lawrence Library",
            "awayTeam":"LUGip",
            "markerImage":"images/red.png",
            "information": "Linux users group meets second Wednesday of each month.",
            "fixture":"Wednesday 7pm",
            "capacity":"",
            "previousScore":""
        },
        {
            "point":"40.211600,-74.695702",
            "homeTeam":"Hamilton Library",
            "awayTeam":"LUGip HW SIG",
            "markerImage":"images/white.png",
            "information": "Linux users can meet the first Tuesday of the month to work out harward and configuration issues.",
            "fixture":"Tuesday 7pm",
            "capacity":"",
            "tv":""
        },
        {
            "point":"40.294535,-74.682012",
            "homeTeam":"Applebees",
            "awayTeam":"After LUPip Mtg Spot",
            "markerImage":"images/newcastle.png",
            "information": "Some of us go there after the main LUGip meeting, drink brews, and talk.",
            "fixture":"Wednesday whenever",
            "capacity":"2 to 4 pints",
            "tv":""
        }
] }`

// From: http://www.sitepoint.com/colors-json-example/
const colors_json_file = `{"markers": [
{
    "colorsArray":[{
            "colorName":"red",
            "hexValue":"#f00"
        },
        {
            "colorName":"green",
            "hexValue":"#0f0"
        },
        {
            "colorName":"blue",
            "hexValue":"#00f"
        },
        {
            "colorName":"cyan",
            "hexValue":"#0ff"
        },
        {
            "colorName":"magenta",
            "hexValue":"#f0f"
        },
        {
            "colorName":"yellow",
            "hexValue":"#ff0"
        },
        {
            "colorName":"black",
            "hexValue":"#000"
}]}]}
`

// From http://www.sitepoint.com/twitter-json-example/
// escaped (backslashed) quotes in source url
const twitter_json = `{"results":[
     {"text":"@twitterapi  http://tinyurl.com/ctrefg",
     "to_user_id":396524,
     "to_user":"TwitterAPI",
     "from_user":"jkoum",
     "metadata":
     {
      "result_type":"popular",
      "recent_retweets": 109
     },
     "id":1478555574,   
     "from_user_id":1833773,
     "iso_language_code":"nl",
     "source":"<a href=\"http://twitter.com/\">twitter< /a>",
     "profile_image_url":"http://s3.amazonaws.com/twitter_production/profile_images/118412707/2522215727_a5f07da155_b_normal.jpg",
     "created_at":"Wed, 08 Apr 2009 19:22:10 +0000"}],

     "since_id":0,
     "max_id":1480307926,
     "refresh_url":"?since_id=1480307926&amp;q=%40twitterapi",
     "results_per_page":15,
     "next_page":"?page=2&amp;max_id=1480307926&amp;q=%40twitterapi",
     "completed_in":0.031704,
     "page":1,
     "query":"%40twitterapi"
}`

// From http://www.sitepoint.com/youtube-json-example/
const youtube_json = `{"apiVersion":"2.0",
 "data":{
    "updated":"2010-01-07T19:58:42.949Z",
    "totalItems":800,
    "startIndex":1,
    "itemsPerPage":1,
    "items":[
        {"id":"hYB0mn5zh2c",
         "uploaded":"2007-06-05T22:07:03.000Z",
         "updated":"2010-01-07T13:26:50.000Z",
         "uploader":"GoogleDeveloperDay",
         "category":"News",
         "title":"Google Developers Day US - Maps API Introduction",
         "description":"Google Maps API Introduction ...",
         "tags":[
            "GDD07","GDD07US","Maps"
         ],
         "thumbnail":{
            "default":"http://i.ytimg.com/vi/hYB0mn5zh2c/default.jpg",
            "hqDefault":"http://i.ytimg.com/vi/hYB0mn5zh2c/hqdefault.jpg"
         },
         "player":{
            "default":"http://www.youtube.com/watch?vu003dhYB0mn5zh2c"
         },
         "content":{
            "1":"rtsp://v5.cache3.c.youtube.com/CiILENy.../0/0/0/video.3gp",
            "5":"http://www.youtube.com/v/hYB0mn5zh2c?f...",
            "6":"rtsp://v1.cache1.c.youtube.com/CiILENy.../0/0/0/video.3gp"
         },
         "duration":2840,
         "aspectRatio":"widescreen",
         "rating":4.63,
         "ratingCount":68,
         "viewCount":220101,
         "favoriteCount":201,
         "commentCount":22,
         "status":{
            "value":"restricted",
            "reason":"limitedSyndication"
         },
         "accessControl":{
            "syndicate":"allowed",
            "commentVote":"allowed",
            "rate":"allowed",
            "list":"allowed",
            "comment":"allowed",
            "embed":"allowed",
            "videoRespond":"moderated"
         }
        }
    ]
 }
}`

var validMaps = []string{smallestValidJsonMap, validJsonMap0, validJsonMap1, google_maps_json_file, colors_json_file, twitter_json, youtube_json}

var validArrays = []string{smallestValidJsonArray, validJsonArray0}
