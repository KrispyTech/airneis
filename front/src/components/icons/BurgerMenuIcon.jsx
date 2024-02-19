import webConfig from "@@/webConfig"

const BurgerMenuIcon = ({ color = webConfig.icons.defaultColor, ...otherProps }) => (
  <svg width="31" height="27" viewBox="0 0 31 27" fill="none" xmlns="http://www.w3.org/2000/svg" {...otherProps}>
    <path
      d="M0.5 19.125H30.5V15.375H0.5V19.125ZM0.5 26.625H30.5V22.875H0.5V26.625ZM0.5 11.625H30.5V7.875H0.5V11.625ZM0.5 0.375V4.125H30.5V0.375H0.5Z"
      fill={color}
    />
  </svg>
)

export default BurgerMenuIcon
