import {Pressable, Text, View} from "react-native"
import screens from "../utils/constants/screens"
import Page from "../components/Page"
import React, {useState} from "react"
import FormField from "../components/FormField"
import {createUsers} from "../utils/services/usersServices"
import {Formik} from "formik"
import {useTranslation} from "react-i18next"

import * as yup from "yup"
import errors from "../utils/constants/errors"

const initialValues = {
    firstName: "",
    lastName: "",
    email: "",
    password: ""
}

const SignUpScreen = props => {
    const [error, setError] = useState("")
    const {t} = useTranslation()
    const {navigation} = props
    const handleSignUp =
        errorsList =>
            async ({firstName, lastName, email, password}) => {
                if (errorsList) {
                    setError(t(errors.signup.invalidForm))

                    return
                }

                try {
                    await createUsers({firstName, lastName, email, password})
                    navigation.navigate(screens.signInScreen)
                } catch (e) {
                    setError(t(errors.signup.invalidForm))
                }
            }

    const validationSchema = yup.object({
        firstName: yup.string().required(t(errors.signup.requiredFirstName)),
        lastName: yup.string().required(t(errors.signup.requiredLastName)),
        email: yup.string().email(t(errors.signup.invalidEmail)).required(t(errors.signup.requiredEmail)),
        password: yup.string().required(t(errors.signup.requiredPassword))
    })

    const formData = [
        {
            name: "firstName",
            label: t("signup.firstName"),
            placeholder: t("signup.firstNamePlaceholder"),
            secure: false
        },
        {
            name: "lastName",
            label: t("signup.lastName"),
            placeholder: t("signup.lastNamePlaceholder"),
            secure: false
        },
        {
            name: "email",
            label: t("signup.email"),
            placeholder: t("signup.emailPlaceholder"),
            secure: false
        },
        {
            name: "password",
            label: t("signup.password"),
            placeholder: t("signup.passwordPlaceholder"),
            secure: true
        }
    ]

    return (
        <Page navigation={navigation}>
            {error && <Text className="text-red-500 font-semibold text-lg text-center py-3">{error}</Text>}
            <Text className="text-center text-2xl font-bold py-5">{t("signup.title")}</Text>
            <Formik initialValues={initialValues} onSubmit={handleSignUp()} validationSchema={validationSchema}>
                {({handleChange, handleBlur, handleSubmit, values, errors: formErrors}) => (
                    <View>
                        {formData.map(({name, label, placeholder}, secure) => (
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
                        <Pressable onPress={handleSubmit}
                                   className="bg-secondary mx-auto w-36 py-2 px-4 rounded-lg mt-5">
                            <Text
                                className="text-white text-center font-semibold text-lg">{t("signup.buttonLabel")}</Text>
                        </Pressable>
                    </View>
                )}
            </Formik>
        </Page>
    )
}

export default SignUpScreen
