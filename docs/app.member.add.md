NAME:
   dpg app member add - Invite users to the specified application space

USAGE:
   dpg app member add [command options] [arguments...]

OPTIONS:
   --token value      [Required] API token
   --app-owner value  [Required] An owner of application(s)
   --app-id value     [Required] An application id. e.g. com.deploygate
   --android          [Required] Either of this or ios flag must be specified
   --ios              [Required] Either of this or android flag must be specified
   --invitees value   [Required] Comma separated names or e-mails of those who you want to invite
   
