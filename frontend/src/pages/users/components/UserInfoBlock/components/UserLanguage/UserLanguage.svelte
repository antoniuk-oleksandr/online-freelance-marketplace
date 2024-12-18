<script lang="ts">
    import {Link} from "svelte-routing";
    import type {Language} from "@/types/Language.ts";

    type UserLanguageProps = {
        language: Language,
        isLast: boolean,
    }

    let {language, isLast}: UserLanguageProps = $props();
    const {id, name} = language;
</script>

<Link to={`/search?language=${id}`}>
    <span class="hover:underline duration-200 ease-out">{name}{isLast ? '' : ','}</span>
</Link>