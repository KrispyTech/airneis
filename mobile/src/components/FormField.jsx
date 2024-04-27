import { Text, TextInput, View } from "react-native"
import clsx from "clsx"

const FormField = props => {
  const { value, setValue, placeholder, label, keyboard = "default", className } = props

  return (
    <View className={clsx(className, "px-5 py-2")}>
      <Text className="font-semibold text-lg">{label}</Text>
      <TextInput
        value={value}
        onChangeText={setValue}
        placeholder={placeholder}
        keyboardType={keyboard}
        className={"border border-black py-2 px-3 rounded-lg w-full mt-2"}
      />
    </View>
  )
}

export default FormField
