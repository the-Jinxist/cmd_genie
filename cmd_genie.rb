# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class CmdGenie < Formula
  desc "CLI app to suggest prompted commands"
  homepage "https://github.com/the-Jinxist/cmd_genie"
  version "0.4.1"

  on_macos do
    on_intel do
      url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.4.1/cmd_genie_Darwin_x86_64.tar.gz"
      sha256 "608ac1f503fe823245bfb3fecc8540fcc92195bd56f7b5146399d43280a9ef47"

      def install
        bin.install "cmd_genie"
      end
    end
    on_arm do
      url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.4.1/cmd_genie_Darwin_arm64.tar.gz"
      sha256 "77f9bea4300b3e248e0b89789e358289dd41238e6d3140bc82a6aa9c307071f9"

      def install
        bin.install "cmd_genie"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.4.1/cmd_genie_Linux_x86_64.tar.gz"
        sha256 "4dc8b7d7166a200d0bc523ab9052da77724c57002838724e13b1941e76251c84"

        def install
          bin.install "cmd_genie"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.4.1/cmd_genie_Linux_arm64.tar.gz"
        sha256 "ff06bfd85e97e5837d0381f1cc1a9cc7089c5eece81d3b21ca86ef03868ddf8d"

        def install
          bin.install "cmd_genie"
        end
      end
    end
  end

  def post_install
    cmd_genie --help
  end
end
