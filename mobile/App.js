import React from "react"
import { NavigationContainer } from "@react-navigation/native"
import { createNativeStackNavigator } from "@react-navigation/native-stack"
import HomeScreen from "./src/screens/HomeScreen"
import CategoriesScreen from "./src/screens/CategoriesScreen"
import screens from "./src/utils/constants/screens"

const { Navigator, Screen } = createNativeStackNavigator()
const App = () => (
  <NavigationContainer>
    <Navigator screenOptions={{ headerShown: false }}>
      <Screen name={screens.homePage} component={HomeScreen} />
      <Screen name={screens.categoryPage} component={CategoriesScreen} />
    </Navigator>
  </NavigationContainer>
)

export default App
