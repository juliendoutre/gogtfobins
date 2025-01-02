# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gogtfobins < Formula
  desc ""
  homepage ""
  version "1.13.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.13.0/gogtfobins_Darwin_x86_64.tar.gz"
      sha256 "14078ab0c173d6c5f2659bae228285d042fd54c76db25af4bbf9ef1c7731ccb0"

      def install
        bin.install "gogtfobins"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.13.0/gogtfobins_Darwin_arm64.tar.gz"
      sha256 "dfd48803bb35e0e609a58f5858e39be45ddc7f124081aa96602f083cd8cd5cd9"

      def install
        bin.install "gogtfobins"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.13.0/gogtfobins_Linux_x86_64.tar.gz"
        sha256 "82802364634f4cfa2347d8e9685bebbb03a3f67b7aded53a8beb63f9d980d8ab"

        def install
          bin.install "gogtfobins"
        end
      end
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.13.0/gogtfobins_Linux_arm64.tar.gz"
        sha256 "cc4fde74e486b81487421431cdac97481c195b9b6fdd716609171cb16d288870"

        def install
          bin.install "gogtfobins"
        end
      end
    end
  end
end
