# \TeamsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvite**](TeamsApi.md#AcceptInvite) | **Post** /teams/accept | Accept invite
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /teams | Create team
[**DeleteInvite**](TeamsApi.md#DeleteInvite) | **Delete** /teams/{teamId}/invites/{code} | Delete invite
[**DeleteMember**](TeamsApi.md#DeleteMember) | **Delete** /teams/{teamId}/members/{userId} | Remove member
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /teams/{teamId} | Remove team
[**GenerateInvite**](TeamsApi.md#GenerateInvite) | **Post** /teams/{teamId}/invites | Generate invite
[**ListInvites**](TeamsApi.md#ListInvites) | **Get** /teams/{teamId}/invites | List invites
[**ListTeams**](TeamsApi.md#ListTeams) | **Get** /teams | List teams
[**RetrieveInvite**](TeamsApi.md#RetrieveInvite) | **Get** /teams/{teamId}/invites/{code} | Retrieve invite
[**RetrieveMember**](TeamsApi.md#RetrieveMember) | **Get** /teams/{teamId}/members/{userId} | Retrieve member
[**RetrieveTeam**](TeamsApi.md#RetrieveTeam) | **Get** /teams/{teamId} | Retrieve team
[**RetrieveTeamMembers**](TeamsApi.md#RetrieveTeamMembers) | **Get** /teams/{teamId}/members | List members
[**UpdateMember**](TeamsApi.md#UpdateMember) | **Patch** /teams/{teamId}/members/{userId} | Update member
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Patch** /teams/{teamId} | Update team



## AcceptInvite

> Team AcceptInvite(ctx, body)

Accept invite

Accept an invite from another user. This will add the currently logged in user to the team as a regular memeber. When the invite is accepted it is removed from the team's invites and cannot be reused.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AcceptInviteRequest**](AcceptInviteRequest.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> Team CreateTeam(ctx, body)

Create team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Team**](Team.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvite

> map[string]interface{} DeleteInvite(ctx, teamId, code)

Delete invite

Delete an invite created earlier. You must be an administrator of the team to perform this action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 
**code** | **string**| The invite code. | 

### Return type

**map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMember

> Member DeleteMember(ctx, teamId, userId)

Remove member

Remove a member from the team. You must be an administrator to do this. You can't remove yourself from the team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 
**userId** | **string**| The user ID | 

### Return type

[**Member**](Member.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> Team DeleteTeam(ctx, teamId)

Remove team

Update the team. You must be an administrator of the team to edit it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 

### Return type

[**Team**](Team.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateInvite

> Invite GenerateInvite(ctx, teamId, body)

Generate invite

Update the team. You must be an administrator of the team to edit it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 
**body** | [**InviteRequest**](InviteRequest.md)|  | 

### Return type

[**Invite**](Invite.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvites

> InviteList ListInvites(ctx, teamId)

List invites

Update the team. You must be an administrator of the team to edit it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 

### Return type

[**InviteList**](InviteList.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeams

> TeamList ListTeams(ctx, )

List teams

Update the team. You must be an administrator of the team to edit it.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TeamList**](TeamList.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveInvite

> Invite RetrieveInvite(ctx, teamId, code)

Retrieve invite

Read a single invite from the team's set of non-redeemed invites.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 
**code** | **string**| The invite code. | 

### Return type

[**Invite**](Invite.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMember

> Member RetrieveMember(ctx, teamId, userId)

Retrieve member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 
**userId** | **string**| The user ID | 

### Return type

[**Member**](Member.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTeam

> Team RetrieveTeam(ctx, teamId)

Retrieve team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 

### Return type

[**Team**](Team.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTeamMembers

> MemberList RetrieveTeamMembers(ctx, teamId)

List members

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**| The team ID | 

### Return type

[**MemberList**](MemberList.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> Member UpdateMember(ctx, teamId, userId, body)

Update member

You must be an administrator of the team to update member settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
**userId** | **string**|  | 
**body** | [**Member**](Member.md)|  | 

### Return type

[**Member**](Member.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> Team UpdateTeam(ctx, teamId, body)

Update team

Update the team. You must be an administrator of the team to edit it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
**body** | [**Team**](Team.md)|  | 

### Return type

[**Team**](Team.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

