import { Text, View } from "react-native"
import { MagnifyingGlassIcon, Bars3Icon, ShoppingCartIcon } from "react-native-heroicons/outline"
import ButtonWithSvg from "./ButtonWithSvg"
import mobileConfig from "../../mobileConfig.js"

const Header = () => (
  <View className="bg-[#382E2B] py-1 flex items-center flex-row justify-between px-4">
    <Text className="text-[#FFEDE6] text-4xl font-extrabold">ÀIRNEIS</Text>
    <View className="flex flex-row items-center" gap="8">
      <ButtonWithSvg>
        <MagnifyingGlassIcon {...mobileConfig.icons.lgIcon} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <ShoppingCartIcon {...mobileConfig.icons.lgIcon} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <Bars3Icon {...mobileConfig.icons.xlIcon} />
      </ButtonWithSvg>
    </View>
  </View>
)

export default Header
