# Access Graph Service
This document provides a high-level overview of the key objects and functionalities within the Access Graph service.

## Core Functionality:

Manages user identities and their associated Decentralized Identifiers (DIDs). (if applicable)
Defines and enforces fine-grained access control policies for:
User access to data stored on the application server.
Service-to-service communication within the application ecosystem.
Facilitates secure communication between user wallets (DID containers), the application server, and other services.

### Objects:

#### User (if applicable):

Represents a registered user within the system. (May be replaced with alternative user identification methods if DIDs are not used)
Attributes: (Depending on implementation)
User ID (unique identifier for the user)
Username (optional)
Additional user attributes relevant for access control policies (optional)

#### DID (Optional):

Represents a Decentralized Identifier associated with a user.
Attributes: (Depending on DID standard)
DID string (unique identifier on the blockchain)
Public key(s) associated with the DID

## Service:

Represents a service within the application ecosystem.

## Attributes:
Service ID (unique identifier for the service)
Service description (optional)
Public key or token for service authentication (optional)

## Resource:

Represents a data resource stored on the application server.

### Attributes:
Resource ID (unique identifier for the resource)
Resource name (descriptive name for the resource)
Access Control Policy (ACP):

Defines access permissions for users or services.
Attributes:
Subject (user ID, DID, or service ID)
Resource (resource ID)
Operation (allowed operations on the resource - read, write, etc.)
Conditions (optional - additional constraints for access permission)
Interactions:

Users (through wallets if applicable) or services present their DID or credentials (tokens) when requesting access to resources or interacting with other services.
The Access Graph verifies the user/service identity and checks the relevant access control policies to determine if access is granted.
The Access Graph communicates authorization decisions (granted/denied) to users/services.
Benefits:

User Control (if applicable): Users have control over their data by managing their DIDs and access permissions.
Secure Access Control: The Access Graph enforces access control policies to ensure only authorized users/services can access resources.
Decentralized Management (if applicable): DIDs provide a decentralized approach to user identity management.
Scalable Communication: The Access Graph facilitates secure communication between different entities within the system.
Note: This is a high-level overview. Specific implementation details may vary depending on your chosen technologies and access control strategies.