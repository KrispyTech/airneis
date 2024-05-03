export const getServerSideProps = async ({ locale }) => {

  return { props: { ...(await serverSideTranslations(locale, ["common"])) } }
}
const LoginPage = () => (
  <Page>
    <div className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
      <Text>Login Page Page</Text>
    </div>
    
  </Page>
)

export default LoginPage
