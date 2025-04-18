<script lang="ts">
  import Sign from '@/common-components/Sign/Sign.svelte'
  import ForgotPasswordInputs from '@/pages/forgot-password/components/ForgotPasswordInputs.svelte'
  import { forgotPasswordSchema, initialForgotPasswordData } from '@/pages/forgot-password/helpers'
  import { postAuthRequest } from '@/api/post-auth-request'

  let showEmailSentMessage = $state(false)
  const setShowEmailSentMessage = (value: boolean) => (showEmailSentMessage = value)

  const postForgotPasswordRequest = (data: any) =>
    postAuthRequest('forgot-password', undefined, data)
</script>

<Sign
  keepSignedIn={false}
  defaultValues={initialForgotPasswordData}
  submitAction={postForgotPasswordRequest}
  signText={'Forgot your password'}
  subSignText={'Please enter your account email address.'}
  signButtonText={'Reset Password'}
  schema={forgotPasswordSchema}
  {showEmailSentMessage}
  {setShowEmailSentMessage}
  showBackButton={true}
>
  <ForgotPasswordInputs />
</Sign>
