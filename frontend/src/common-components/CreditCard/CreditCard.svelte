<script lang="ts">
  import { getFileServerHost } from '@/utils/utils'
  import { themeStore } from '@/common-stores/theme-storage'
  import CreditCardLayout from '@/common-components/CreditCard/CreditCardLayout.svelte'
  import CreditCardFrontSide from '@/common-components/CreditCard/components/CreditCardFrontSide/CreditCardFrontSide.svelte'
  import CreditCardBackSide from '@/common-components/CreditCard/components/CreditCardBackSide/CreditCardBackSide.svelte'

  type CreditCardProps = {
    flip: boolean
    securityCode: string
    cardNumber: string
    cardHolderName: string
    expiryDate: string
  }

  const props: CreditCardProps = $props()

  const host = getFileServerHost()
  const mastercardLogo = `${host}/credit-card-logo.png`
  const chip = `${host}/credit-card-chip.png`

  let cardBg = $state('')

  themeStore.subscribe((value) => {
    cardBg = value ? `${host}/credit-card-bg.png` : `${host}/credit-card-light-bg.png`
  })
</script>

<CreditCardLayout flip={props.flip}>
  <CreditCardFrontSide {...props} {cardBg} {chip} {mastercardLogo} />
  <CreditCardBackSide {...props} {cardBg} />
</CreditCardLayout>
