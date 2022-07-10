# System Integrity Go SDK project

## A hobby project where a developer-friendly, efficient, and effective tool for verifying system integrity of (sensitive) devices Go SDK will be developed

The goal of this project is to create an SDK that enables Go developers to easily implement the maintaining of consistency, accuracy and trustworthyness of system data during (physical) transit in their products or services. 

**Go developers can use the SDK to:**

* Generate system data hash
* Generate system data integrity certificate 
* Generate keypairs
* Verify system data hash
* Verify system data integrity certificate

**SDK limitations**

***Disclaimer: This is a hobby-project and should not be used in a production environment***

* Only the following hashing algorithms will be implemented first for the time being:
  * SHA3-256
  * SHA3-384
  * SHA3-512

  If you are looking for other hashing algorithms, please stick around. More hashing alogithms are intended to be implemented in the future.

* Only ECDSA will be implemented for the time being, and only the following curves:
  * NIST P-256
  * NIST P-384
  * NIST P-521

  If you are looking for other EC-curves or even RSA, please stick around. More curves and DSA are intented to be implemented in the future.

## Documentation

You can find documentation on Readthedocs: https://system-integrity-go-sdk.readthedocs.io/en/latest/

## Visuals

UML-diagrams are coming soon

## Installation instructions

1. Open your terminal
2. Type (or copy and paste) ``` go install https://github.com/SeanVisser1998/System-Integrity-Go-SDK ```
3. Run the command

## User instructions

User instructions are coming soon

## Known issues

There are currently no known issues with this project

## Found a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above. If you would like to submit a pull request with a fix, please reference the issue you created. When submiting a pull request, please make sure to follow the coding style guidelines.

## SDK development guidelines

**Coding style guidelines**

* Only use Go Standard Library in order to avoid suprises when dependancies are updated or updated and to limit vulnerabilities inherited from dependancies
* Make configuration simple in order to enhance developer-friendliness
* Do not authenticate within the SDK, this is a responsibility that lies with the application
* Write code according to the idiomatic Go style to ensure effective Go-code. https://go.dev/doc/effective_go
* When naming variables use Go-guidelines for spelling to enhance developer-friendliness. https://github.com/golang/go/wiki/Spelling
* Make use of comments to enhance developer-friendliness. https://go.dev/doc/comment
* Handle errors according to Go-guidelines to enhance developer-friendliness. https://go.dev/blog/error-handling-and-go
* Auto generate tests to ensure high test-coverage and limit time spend on writing test-cases. https://github.com/cweill/gotests
* More coming soon...

**Versioning**

The SDK will follow semantic versioning guidelines. More information about semantic versioning can be found at: https://semver.org

