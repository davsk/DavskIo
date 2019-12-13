// copyright

// Package DavskIo provides a common I/O interface for GUI, console, web, and batch for Davsk apps.
// 
// Overview
//
// Apps created by Davsk may use DI (Dependency Injection) to allow the user to select I/O options.
// The DavskIo interface defines the methods implemented.
//
// Instructions
//
// Allow the User to config the app to select their preferred I/O method.
// Avoid using any color codes ever, instead define the class and allow implementation to be user and hardware specific.
//     DavskIo.MsgBox(ClassError,"Heading","Error Message!",BtnOk)
//
// Legal
//
// Copyright (c) 2019 Davsk Ltd. Co.
// All Rights Reserved. Licensed MIT.
package DavskIo // import "pkg.davsk.net/DavskIo"
