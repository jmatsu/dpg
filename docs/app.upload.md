NAME:
   dpg app upload - Upload either android application or iOS application to the specified owner space

USAGE:
   dpg app upload [command options] [arguments...]

OPTIONS:
   --token value             [Required] API token
   --app-owner value         [Required] An owner of application(s)
   --android                 [Required] Either of this or ios flag must be specified
   --ios                     [Required] Either of this or android flag must be specified
   --app value               [Required] A path of an application file to be uploaded
   --public                  Specify true if an application to be uploaded should be public
   --enableNotification      [iOS only] Specify true if iOS's notifications should be enabled
   --message value           A short message to explain this update
   --distributionKey value   A key of a distribution which an application will be uploaded to
   --distributionName value  A name of a distribution which an application will be uploaded to
   --releaseNote value       A release note for this revision
   
