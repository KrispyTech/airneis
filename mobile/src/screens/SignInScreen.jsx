import { Pressable, Text, View } from "react-native"
import Page from "../components/Page"
import React, { useState } from "react"
import { useTranslation } from "react-i18next"
import errors from "../utils/constants/errors"
import * as yup from "yup"
import { passwordRegex } from "../utils/constants/password"
import { Formik } from "formik"
import FormField from "../components/FormField"
import { createSession } from "../utils/services/sessionServices"

const initialValues = {
  email: "",
  password: ""
}

const SignInScreen = props => {
  const [error, setError] = useState("")
  const { t } = useTranslation()
  const { navigation } = props
  const handleSignUp =
    errorsList =>
    async ({ email, password }) => {
      if (errorsList) {
        setError(t(errors.signin.invalidForm))

        return
      }

      try {
        await createSession({ email, password })
      } catch (e) {
        setError(t(errors.signin.invalidForm))
      }
    }

  const validationSchema = yup.object({
    email: yup.string().email(t(errors.signup.invalidEmail)).required(t(errors.signup.requiredEmail)),
    password: yup.string().required(t(errors.signup.requiredPassword)).matches(passwordRegex, t(errors.signup.invalidPassword))
  })

  const formData = [
    {
      name: "email",
      label: t("signIn.email"),
      placeholder: t("signIn.emailPlaceholder"),
      secure: false
    },
    {
      name: "password",
      label: t("signIn.password"),
      placeholder: t("signIn.passwordPlaceholder"),
      secure: true
    }
  ]

  return (
    <Page navigation={navigation}>
      {error && <Text className="text-red-500 font-semibold text-lg text-center py-3">{error}</Text>}
      <Text className="text-center text-2xl font-bold py-5">{t("signIn.title")}</Text>
      <Formik initialValues={initialValues} onSubmit={handleSignUp()} validationSchema={validationSchema}>
        {({ handleChange, handleBlur, handleSubmit, values, errors: formErrors }) => (
          <View>
            {formData.map(({ name, label, placeholder, secure }) => (
              <FormField
                key={name}
                value={values[name]}
                setValue={handleChange(name)}
                onBlur={handleBlur(name)}
                label={label}
                placeholder={placeholder}
                error={formErrors[name]}
                secure={secure}
              />
            ))}
            <Pressable onPress={handleSubmit} className="bg-secondary mx-auto w-36 py-2 px-4 rounded-lg mt-5">
              <Text className="text-white text-center font-semibold text-lg">{t("signIn.buttonLabel")}</Text>
            </Pressable>
          </View>
        )}
      </Formik>
    </Page>
  )
}

export default SignInScreen
