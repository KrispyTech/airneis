import { Text, View } from "react-native"
import { MagnifyingGlassIcon, Bars3Icon, ShoppingCartIcon } from "react-native-heroicons/outline"
import ButtonWithSvg from "./ButtonWithSvg"
import mobileConfig from "../../mobileConfig"
import { colors } from "../styles"

const Header = () => (
  <View className="bg-secondary py-1 flex items-center flex-row justify-between px-4">
    <Text className="text-primary text-4xl font-extrabold">ÀIRNEIS</Text>
    <View className="flex flex-row items-center">
      <ButtonWithSvg>
        <MagnifyingGlassIcon {...mobileConfig.icons.lgIcon} color={colors.primary} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <ShoppingCartIcon {...mobileConfig.icons.lgIcon} color={colors.primary} style={{ marginLeft: 8 }} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <Bars3Icon {...mobileConfig.icons.xlIcon} color={colors.primary} style={{ marginLeft: 8 }} />
      </ButtonWithSvg>
    </View>
  </View>
)

export default Header
