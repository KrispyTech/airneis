import { Text, View } from "react-native"
import { MagnifyingGlassIcon, Bars3Icon, ShoppingCartIcon } from "react-native-heroicons/outline"
import ButtonWithSvg from "./ButtonWithSvg"
import { color, icons } from "../../tailwind.config.cjs"

const Header = () => (
  <View className={`bg-[${color.brown}] py-1 flex items-center flex-row justify-between px-4`}>
    <Text className={`text-[${color.lightPink}] text-4xl font-extrabold`}>ÀIRNEIS</Text>
    <View className="flex flex-row items-center">
      <ButtonWithSvg>
        <MagnifyingGlassIcon {...icons.lgIcon} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <ShoppingCartIcon {...icons.lgIcon} style={{ marginLeft: 8 }} />
      </ButtonWithSvg>
      <ButtonWithSvg>
        <Bars3Icon {...icons.xlIcon} style={{ marginLeft: 8 }} />
      </ButtonWithSvg>
    </View>
  </View>
)

export default Header
