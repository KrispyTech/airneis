import React from "react"
import { NavigationContainer } from "@react-navigation/native"
import { createNativeStackNavigator } from "@react-navigation/native-stack"
import HomeScreen from "./src/screens/HomeScreen"
import CategoryScreen from "./src/screens/CategoryScreen"

import screens from "./src/utils/constants/screens"

const { Navigator, Screen } = createNativeStackNavigator()
const App = () => (
  <NavigationContainer>
    <Navigator screenOptions={{ headerShown: false }}>
      <Screen name={screens.homePage} component={HomeScreen} />
      <Screen name={screens.categoryPage} component={CategoryScreen} />
    </Navigator>
  </NavigationContainer>
)

export default App
