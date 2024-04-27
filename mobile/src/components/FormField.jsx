import {Text, TextInput, View} from "react-native"
import clsx from "clsx"

const FormField = props => {
    const {value, setValue, placeholder, label, onBlur, error, className, secure = false} = props

    return (
        <View className={clsx(className, "px-5 py-2")}>
            <Text className="font-semibold text-lg">{label}</Text>
            <TextInput
                value={value}
                onChangeText={setValue}
                placeholder={placeholder}
                onBlur={onBlur}
                secureTextEntry={secure}
                className={"border border-black py-2 px-3 rounded-lg w-full mt-2"}
            />
            <Text className="text-red-500 mt-2 font-semibold">{error}</Text>
        </View>
    )
}

export default FormField
