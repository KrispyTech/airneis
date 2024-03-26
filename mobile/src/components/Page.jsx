import { SafeAreaView, ScrollView, View } from "react-native"
import Header from "./Header"

const Page = props => {
  const { children, navigation } = props

  return (
    <View className="h-screen flex flex-col w-full h-screen">
      <SafeAreaView className="bg-secondary">
        <Header navigation={navigation} />
        <ScrollView className="grow overflow-y-auto flex flex-col w-full h-full bg-primary">{children}</ScrollView>
      </SafeAreaView>
    </View>
  )
}

export default Page
