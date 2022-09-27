$version: "2"
namespace example.Person
use aws.protocols#awsJson1_1

enum PhoneType {
    MOBILE
    HOME
    WORK
}

structure PhoneNumber {
  number: String
  type: PhoneType
}

list PhoneNumbers {
    member: PhoneNumber
}

/// Provides weather forecasts.
structure Person {
    @required
    name: String
    @required
    birthdate: Timestamp
    @required
    email: String
    numbers: PhoneNumbers
}
