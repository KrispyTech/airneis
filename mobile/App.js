import React from "react"
import { NavigationContainer } from "@react-navigation/native"
import { createNativeStackNavigator } from "@react-navigation/native-stack"
import HomeScreen from "./src/screens/HomeScreen"
import CategoryScreen from "./src/screens/CategoryScreen"

import screens from "./src/utils/constants/screens"
import ProductScreen from "./src/screens/ProductScreen"
import SignUpScreen from "./src/screens/SignUpScreen"

const { Navigator, Screen } = createNativeStackNavigator()
const App = () => (
  <NavigationContainer>
    <Navigator screenOptions={{ headerShown: false }}>
      <Screen name={screens.homePage} component={HomeScreen} />
      <Screen name={screens.categoryPage} component={CategoryScreen} />
      <Screen name={screens.productPage} component={ProductScreen} />
      <Screen name={screens.signUpScreen} component={SignUpScreen} />
    </Navigator>
  </NavigationContainer>
)

export default App
