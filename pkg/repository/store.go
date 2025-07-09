package repository

const (
	JsonCampaigns = `
[
    {
        "campaignId": "spotify",
        "campaignName": "Spotify - Music for everyone",
        "creatives": [
            {
                "imageUrl": "https://somelink"
            }
        ],
        "cta": "Download",
        "campaignStatus": "ACTIVE"
    },
    {
        "campaignId": "duolingo",
        "campaignName": "Duolingo: Best way to learn",
        "creatives": [
            {
                "imageUrl": "https://somelink2"
            }
        ],
        "cta": "Install",
        "campaignStatus": "ATCIVE"
    },
    {
        "campaignId": "subwaysurfer",
        "campaignName": "Subway Surfer",
        "creatives": [
            {
                "imageUrl": "https://somelink3"
            }
        ],
        "cta": "Play",
        "campaignStatus": "ATCIVE"
    },
    {
        "campaignId": "inactiveCampaign",
        "campaignName": "Inactive Test Campaign",
        "creatives": [
            {
                "imageUrl": "https://somelink4"
            }
        ],
        "cta": "Buy Now",
        "campaignStatus": "INATCIVE"
    },
    {
        "campaignId": "netflixPromo",
        "campaignName": "Netflix: Stream your favorite shows",
        "creatives": [
            {
                "imageUrl": "https://somelink_netflix_1"
            },
            {
                "imageUrl": "https://somelink_netflix_2"
            }
        ],
        "cta": "Sign Up",
        "campaignStatus": "ACTIVE"
    },
    {
        "campaignId": "amazonPrimeDay",
        "campaignName": "Amazon Prime Day Deals",
        "creatives": [
            {
                "imageUrl": "https://somelink_amazon_1"
            }
        ],
        "cta": "Shop Now",
        "campaignStatus": "ACTIVE"
    },
    {
        "campaignId": "googleAdsTutorial",
        "campaignName": "Learn Google Ads in 30 Days",
        "creatives": [
            {
                "imageUrl": "https://somelink_google_1"
            },
            {
                "imageUrl": "https://somelink_google_2"
            },
            {
                "imageUrl": "https://somelink_google_3"
            }
        ],
        "cta": "Start Free Trial",
        "campaignStatus": "INACTIVE"
    }
]`

	JsonRules = `
[
    {
        "campaignId": "spotify",
        "includeCountries": [
            "US",
            "Canada"
        ]
    },
    {
        "campaignId": "duolingo",
        "includeOs": [
            "Android",
            "iOS"
        ],
        "excludeCountries": [
            "US"
        ]
    },
    {
        "campaignId": "subwaysurfer",
        "includeOs": [
            "Android"
        ],
        "includeAppIDs": [
            "com.gametion.ludokinggame"
        ]
    },
    {
        "campaignId": "netflixPromo",
        "includeCountries": [
            "GB",
            "DE",
            "FR"
        ],
        "excludeOs": [
            "Windows"
        ]
    },
    {
        "campaignId": "amazonPrimeDay",
        "includeCountries": [
            "US",
            "UK",
            "JP"
        ],
        "includeOs": [
            "Android",
            "iOS",
            "Web"
        ]
    },
    {
        "campaignId": "googleAdsTutorial",
        "includeOs": [
            "Web",
			"iOS"
        ],
        "excludeCountries": [
            "CN",
            "RU"
        ]
    },
    {
        "campaignId": "tiktokChallenge",
        "includeCountries": [
            "ID",
            "PH",
            "VN"
        ],
        "includeOs": [
            "iOS"
        ]
    },
    {
        "campaignId": "mcdonaldsAppPromo",
        "includeCountries": [
            "US"
        ],
        "includeOs": [
            "Android",
            "iOS"
        ],
        "excludeAppIDs": [
            "com.burgerking.app"
        ]
    },
	{
	"campaignId": "testcampaignstargetengine",
	"includeOs" : [
		"iOS"
	]
	}
]
`
)
