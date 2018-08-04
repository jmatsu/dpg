NAME:
   dpg app upload - Upload either android application or iOS application to the specified owner space

USAGE:
   dpg app upload [command options] [arguments...]

OPTIONS:
   --token value             [Required] API token
   --app-owner value         [Required] The owner of the application
   --android                 [Required] Specify this if the application is an android application
   --ios                     [Required] Specify this if the application is an iOS application
   --app value               [Required] The file path of the application to be uploaded
   --public                  Specify true if an application to be uploaded should be public
   --enableNotification      [iOS only] Specify true if iOS's notifications should be enabled
   --message value           A short message to explain this update
   --distributionKey value   A key of a distribution to be updated
   --distributionName value  A name of a distribution to be updated
   --releaseNote value       A release note for this revision
   
