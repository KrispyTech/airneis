import { TouchableHighlight, View } from "react-native"

const ButtonWithSvg = props => {
  const { onPress, children } = props

  return (
    <TouchableHighlight onPress={onPress}>
      <View>{children}</View>
    </TouchableHighlight>
  )
}

export default ButtonWithSvg
