***This project is currently under development***
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

***This project is currently under development***