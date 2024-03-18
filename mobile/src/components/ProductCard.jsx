import { Text, View, Image } from "react-native"
import mobileConfig from "../../mobileConfig"

const ProductCard = props => {
  const {
    product: { name, thumbnailURL, priceWithTaxes }
  } = props

  return (
    <View>
      <Image
        style={mobileConfig.images.fullScreenImageMd}
        className="rounded box-shadow-xl border "
        source={{
          uri: thumbnailURL
        }}
      />
      <View className="flex flex-row justify-between">
        <Text className="text-lg font-light">{name} </Text>
        <Text className="text-lg font-light">{priceWithTaxes} €</Text>
      </View>
    </View>
  )
}

export default ProductCard
