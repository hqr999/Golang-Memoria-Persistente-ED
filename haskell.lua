return {
  -- Haskell syntax highlighting
  {
    "neovimhaskell/haskell-vim",
    ft = "haskell",
  },
  -- Simple Haskell LSP support
  {
    "neovim/nvim-lspconfig",
    opts = {
      servers = {
        hls = {
          filetypes = { "haskell", "lhaskell" },
        },
      },
    },
  },
}
