# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class CmdGenie < Formula
  desc "CLI app to suggest prompted commands"
  homepage "https://github.com/the-Jinxist/cmd_genie"
  version "0.3.3"

  on_macos do
    on_intel do
      url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.3.3/cmd_genie_Darwin_x86_64.tar.gz"
      sha256 "5929fc34a9252db64536b011484b5b5e18df0cebbd606930e6f6381ec64809f9"

      def install
        bin.install "cmd_genie"
      end
    end
    on_arm do
      url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.3.3/cmd_genie_Darwin_arm64.tar.gz"
      sha256 "8b1a0264e5d0c74548ee5fe9e12bde6b8b1e4fae92169a3364f6b0b1ffacae16"

      def install
        bin.install "cmd_genie"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.3.3/cmd_genie_Linux_x86_64.tar.gz"
        sha256 "953cbe200ec455e9818348ac8fcdbb0fc98a8b9ff87637929d08f9a4f98d06ad"

        def install
          bin.install "cmd_genie"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/the-Jinxist/cmd_genie/releases/download/v0.3.3/cmd_genie_Linux_arm64.tar.gz"
        sha256 "e18bf427ffa7a636a71b37c55acf1572bdfb69c3d9e9e25b3688444027328465"

        def install
          bin.install "cmd_genie"
        end
      end
    end
  end
end
