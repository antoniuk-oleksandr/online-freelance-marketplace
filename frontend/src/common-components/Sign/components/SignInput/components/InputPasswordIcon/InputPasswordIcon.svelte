<script lang="ts">
  import Icon from '@iconify/svelte'
  import InputPasswordIconLayout from '@/common-components/Sign/components/SignInput/components/InputPasswordIcon/InputPasswordIconLayout.svelte'

  type InputPasswordIconProps = {
    type: 'text' | 'password'
    showPassword: boolean
    setShowPassword: (value: boolean) => void
  }

  const { type, showPassword, setShowPassword }: InputPasswordIconProps = $props()
</script>

{#if type === 'password'}
  <InputPasswordIconLayout {showPassword} {setShowPassword}>
    {#if showPassword}
      <Icon icon="mdi:eye-off-outline" />
    {:else}
      <Icon icon="mdi:eye-outline" />
    {/if}
  </InputPasswordIconLayout>
{/if}
