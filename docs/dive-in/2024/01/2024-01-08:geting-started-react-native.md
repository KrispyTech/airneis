# REACT NATIVE with EXPO

I did React Native installation with Expo, this library comes with some features that makes you can preview app on your phone or on simulators like: Xcode simulator (ios) and AndroidStudio.

FIRST THING RUN `npm i` !! and then you can `npm start`.

> #### When you start with `npm start` you have access to some commands:
>
> › Press s │ switch to development build  
> › Press a │ open Android  
> › shift+a │ select a device or emulator  
> › Press i │ open iOS simulator (it will open Xcode with a device already specified)  
> › shift+i │ select a simulator (it will ask you to choose a device and keep it in Xcode for the next time you do the command `i`)  
> › Press w │ open web
>
> › Press r │ reload app  
> › Press j │ open debugger  
> › Press m │ toggle menu  
> › shift+m │ more tools  
> › Press o │ open project code in your editor  
> › Press c │ show project QR

If you want the preview on your phone, you have to connect it to your pc for the first use and download `Expo Go` and then if nothing appears when you do `npm start` with phone connected to pc you can try by scaning the QR Code.

I installed nativewind and the doc is at: https://www.nativewind.dev/  
_Not all is explained for all the times where it's not specific they send you to tailwindcss doc._  
To run nativewind properly I had to install a specific tailwindcss version: `npm install tailwindcss@3.3.2 --save-dev`

## Create an ios simulator

If you want to test the app on a simulator instead of your phone you can do it with Xcode simulators.
First download Xcode on the AppStore then press cmd + , to open settings the go to platform and download the version of ios.
The press cmd + space to search for simulator and open it.
Then you can open the app with `npm start` and press `i` to open the ios simulator.
