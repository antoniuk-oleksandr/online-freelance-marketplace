<script lang="ts">
  import Sign from '@/common-components/Sign/Sign.svelte'
  import { initialSignUpData, postSignUpRequest, signUpSchema } from '@/pages/sign-up/helpers'
  import SignUpInputs from '@/pages/sign-up/components/SignUpInputs/SignUpInputs.svelte'

  let showEmailSentMessage = $state(false)
  const setShowEmailSentMessage = (value: boolean) => (showEmailSentMessage = value)
</script>

<Sign
  {showEmailSentMessage}
  {setShowEmailSentMessage}
  defaultValues={initialSignUpData}
  googleButtonText={'Continue with Google'}
  submitAction={postSignUpRequest}
  signText={'Sign Up'}
  subSignText={'Please sign Up to continue to your account.'}
  subFormText={'Already have an account?'}
  signButtonText={'Sign Up'}
  subFormLink={'/sign-in'}
  signButtonLinkText={'Sign In'}
  schema={signUpSchema}
  keepSignedIn={false}
>
  <SignUpInputs />
</Sign>
