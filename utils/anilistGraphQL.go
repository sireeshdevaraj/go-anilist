package utils

const Query = `
	query{
  		User(id: %d){
    	id
    	name
    	bannerImage
    	updatedAt
    	statistics{
      		anime{
				count
				minutesWatched
				episodesWatched
      		}
    	}
  		}
  	
  
  	MediaListCollection(userId:%d,type:ANIME,status_in:[COMPLETED,CURRENT],sort:FINISHED_ON_DESC){
		lists{
      		entries{
				status
				progress
				media{
					title {
						romaji
						english
						native
						userPreferred
						}
            coverImage{
              large
            }
					}
      			}
    		}
  		}
  
}`

/*
query{
  User(id: 5631742){
    id
    name
    bannerImage
    updatedAt
    statistics{
      anime{
        count
        minutesWatched
        episodesWatched
      }
    }
  }
  Activity(userId:5631742,type: MEDIA_LIST){

      ... on ListActivity{
  			progress
        status
        media{
          title {
            romaji
            english
            native
            userPreferred

          }
          coverImage{
            medium
          }
        }
      }

  }

  MediaListCollection(userId:5631742,type:ANIME,status_in:[COMPLETED,CURRENT],sort:FINISHED_ON_DESC){
		lists{
      entries{
        status
        progress
        media{
          title {
            romaji
            english
            native
            userPreferred
          }

        }
      }
    }
  }

}
*/
