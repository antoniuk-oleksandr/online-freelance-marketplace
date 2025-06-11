<script lang="ts">
  import { Link } from 'svelte-routing'

  type AppLogoProps = {
    styles?: String
  }

  const { styles }: AppLogoProps = $props()
</script>

<Link class={styles} to={'/'}>
  <h1 class={'text-2xl text-cyan-500 font-semibold'}>OFM</h1>
</Link>
