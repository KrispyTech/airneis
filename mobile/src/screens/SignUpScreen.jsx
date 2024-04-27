import { Button, Pressable, Text, TextInput, View } from "react-native"
import screens from "../utils/constants/screens"
import Page from "../components/Page"
import React, { useState } from "react"
import FormField from "../components/FormField"
import { createUsers } from "../utils/services/usersServices"

const SignUpScreen = props => {
  const [firstName, setFirstName] = useState("")
  const [lastName, setLastName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")
  const { navigation } = props
  const handleSignUp = () => async () => {
    try {
      console.log({ firstName, lastName, email, password })
      await createUsers({ firstName, lastName, email, password })
    } catch (e) {
      setError(e.message)
      console.log(e)
    }
  }

  return (
    <Page navigation={navigation}>
      {error && <Text className="text-red-500 font-semibold text-lg text-center py-3">{error}</Text>}
      <Text className="text-center text-2xl font-bold py-5">Welcome to Airneis</Text>
      <View>
        <FormField value={firstName} setValue={setFirstName} label="First name" placeholder="Your first name" />
        <FormField value={lastName} setValue={setLastName} label="Last name" placeholder="Your last name" />
        <FormField value={email} setValue={setEmail} label="Email" placeholder="Your email" />
        <FormField value={password} setValue={setPassword} label="Your password" placeholder="Your password" />
        <Pressable onPress={handleSignUp()} className="bg-secondary mx-auto w-36 py-2 px-4 rounded-lg mt-5">
          <Text className="text-white text-center font-semibold text-lg">Sign up</Text>
        </Pressable>
      </View>
    </Page>
  )
}

export default SignUpScreen
