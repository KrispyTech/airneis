import { SafeAreaView, ScrollView, View } from "react-native"
import Header from "./Header"

const Page = props => {
  const { children, navigation } = props

  return (
    <View className="h-screen flex flex-col w-full">
      <SafeAreaView className="bg-secondary"></SafeAreaView>
      <Header navigation={navigation} />
      <ScrollView className="grow overflow-y-auto flex flex-col w-full bg-primary">{children}</ScrollView>
    </View>
  )
}

export default Page
