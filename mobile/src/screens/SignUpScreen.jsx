import {Pressable, Text, View} from "react-native"
import screens from "../utils/constants/screens"
import Page from "../components/Page"
import React, {useState, useTransition} from "react"
import FormField from "../components/FormField"
import {createUsers} from "../utils/services/usersServices"
import {Formik} from "formik"

import * as yup from "yup"
import errors from "../utils/constants/errors"

const initialValues = {
    firstName: "",
    lastName: "",
    email: "",
    password: ""
}

const validationSchema = yup.object(
    {
        firstName: yup.string().required("First name is required"),
        lastName: yup.string().required("Last name is required"),
        email: yup.string().email("Invalid email").required("Email is required"),
        password: yup.string().required("Password is required")
    }
)

const SignUpScreen = props => {
    const [error, setError] = useState("")
    const {t} = useTransition()
    const {navigation} = props
    const handleSignUp = (errorsList) => async ({firstName, lastName, email, password}) => {
        if (errorsList) {
            setError(t(errors.signup.invalidForm))

            return
        }

        try {
            await createUsers({firstName, lastName, email, password})
            navigation.navigate(screens.signInScreen)
        } catch (e) {
            setError(e.message)
        }
    }

    return (
        <Page navigation={navigation}>
            {error && <Text className="text-red-500 font-semibold text-lg text-center py-3">{error}</Text>}
            <Text className="text-center text-2xl font-bold py-5">Welcome to Airneis</Text>
            <Formik initialValues={initialValues} onSubmit={handleSignUp()} validationSchema={validationSchema}>
                {({handleChange, handleBlur, handleSubmit, values, errors}) => (
                    <View>
                        <FormField
                            value={values.firstName}
                            setValue={handleChange("firstName")}
                            onBlur={handleBlur("firstName")}
                            label="First name"
                            name="firstName"
                            placeholder="Your first name"
                            error={errors.firstName}
                        />
                        <FormField
                            value={values.lastName}
                            setValue={handleChange("lastName")}
                            onBlur={handleBlur("lastName")}
                            label="Last name"
                            placeholder="Your last name"
                        />
                        <FormField value={values.email} setValue={handleChange("email")} label="Email"
                                   onBlur={handleBlur("email")} placeholder="Your email"/>
                        <FormField
                            value={values.password}
                            setValue={handleChange("password")}
                            label="Your password"
                            onBlur={handleBlur("password")}
                            placeholder="Your password"
                        />
                        <Pressable onPress={handleSubmit}
                                   className="bg-secondary mx-auto w-36 py-2 px-4 rounded-lg mt-5">
                            <Text className="text-white text-center font-semibold text-lg">Sign up</Text>
                        </Pressable>
                    </View>
                )}
            </Formik>
        </Page>
    )
}

export default SignUpScreen
