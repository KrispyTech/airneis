## Getting started

First install your modules:

```
npm i
```

Then you can start the app:

```
npm start
```

The following menu will be displayed:

> › Press s │ switch to development build <br>
> › Press a │ open Android<br>
> › shift+a │ select a device or emulator<br>
> › Press i │ open iOS simulator (it will open Xcode with a device already specified)<br>
> › shift+i │ select a simulator (it will ask you to choose a device and keep it in Xcode for the next time you do the command i)<br>
> › Press w │ open web
>
> › Press r │ reload app<br>
> › Press j │ open debugger<br>
> › Press m │ toggle menu<br>
> › shift+m │ more tools<br>
> › Press o │ open project code in your editor<br>
> › Press c │ show project QR<br>

If you don't have a virtualization environment yet, you won't be able to visualize the app. The following instructions will help you setup a virtual mobile environment.

## Android

- Install Android Studio: https://developer.android.com/studio<br>
- Follow the setup wizard. It will install a virtual device by default<br>

Press A on the previous menu to start the virtual device automatically

## iOs

- Download Xcode from the AppStore
- Open Settings
- Go to 'Platform' and download your version of iOs
- Press `cmd + space` to search for simulator and open it.

Press i on the previous menu to open the iOs simulator.

## From your mobile

If you want the preview on your phone, you have to connect it to your PC for the first use and download Expo Go.

If nothing appears when you do `npm start` with your phone connected to your PC, you can try by scaning the QR Code.
