# Rating Service

Clone this project
- `git clone https://github.com/abhinavmaury/rating-service.git`

Import in IDE and run-
It will run on port 7000

This service is using MySQL-
- Create table for rating  
 ``CREATE TABLE `user_rating` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(45) NOT NULL,
  `user_type` enum('driver','passenger') NOT NULL,
  `rating_value` tinyint(3) NOT NULL,
  `rating_base` tinyint(3) NOT NULL DEFAULT '5',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);``

## Objective
1. The passenger should be able to rate a given ride.
2. The driver should be able to see aggregated rating of his all rides
3. The driver should be able to rate the passenger after ride
4. The passenger should be able to see his aggregated rating based on all the rides he has taken.

It contains 4 POST REST APIs
- To get Passenger Aggregated Rating - need to input userID  
**Request URL-  {{url}}/api/rating/passenger  
    Request Body- 
    {
    "userID":123
    }**
    
- Rating from Passenger- need to input 'userID', 'rideID', 'rating', 'rating base'  
**Request URL- {{url}}/api/rating/passenger/ride**  
**Request Body-
{
    "userID":123,
    "rideID":3456,
    "rating":90,
    "rating_base":5
}**

- To get Driver Aggregated Rating - need to input userID  
 **Request URL-  {{url}}/api/rating/driver  
Request Body- 
{
 "userID":331
}**
    
- Rating from Driver- need to input 'userID', 'rideID', 'rating', 'rating base'  
**Request URL- {{url}}/api/rating/driver/ride  
Request Body- 
{
    "userID":331,
    "rideID":3462,
    "rating":90,
    "rating_base":5
}**

