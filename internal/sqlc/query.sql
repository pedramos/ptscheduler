-- name: NewTrainee :one
Insert into trainee (name, perweek, late) values (?,?,?) Returning id;

-- name: NewSession :exec
Insert into schedule (traineeid, date) values(?,?);

-- name: AddAvailability :exec
Insert into availability (traineeid, startdate, enddate) values(?,?,?);

-- name: NewUsername :exec
Insert into usernames (traineeid, username, password) values(?,?, ?);

-- name: GetPassword :one
Select password from usernames where username = ?;

-- name: RealName :one
Select name from trainee, usernames where trainee.id = usernames.traineeid and usernames.username = ?;

-- name: ScheduledTrainning :many
Select date from schedule, usernames where schedule.traineeid = usernames.traineeid and usernames.username = ? and schedule.date > date('now', '-1 day') ;

-- name: ScheduledAvailability :many
Select startdate, enddate from availability, usernames where availability.traineeid = usernames.traineeid and usernames.username = ? and availability.enddate > date('now', '-1 day') ;
