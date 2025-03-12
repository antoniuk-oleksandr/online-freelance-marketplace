<script lang="ts">
  type InputElementProps = {
    setIsFocused: (value: boolean) => void
    id: string
    type: 'text' | 'password'
    showPassword: boolean
  }

  const { id, type, setIsFocused, showPassword }: InputElementProps = $props()
</script>

<input
  onfocus={() => setIsFocused(true)}
  onblur={() => setIsFocused(false)}
  class="px-4 pr-9 py-3 w-full outline-none bg-transparent text-base text-light-palette-text-primary dark:text-dark-palette-text-primary"
  {id}
  name={id}
  type={type === 'password' ? (showPassword ? 'text' : 'password') : type}
/>
