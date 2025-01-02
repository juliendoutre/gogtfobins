# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gogtfobins < Formula
  desc ""
  homepage ""
  version "1.14.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.14.0/gogtfobins_Darwin_x86_64.tar.gz"
      sha256 "d6952b7f74212db0a61142b6dba8e42eb5c728885df47bba1a8b184b38e4dfd4"

      def install
        bin.install "gogtfobins"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.14.0/gogtfobins_Darwin_arm64.tar.gz"
      sha256 "b3274973a9857e6d9a3fcf4ce8eec78fa93662d67bff7b056c49dbb2060b9a79"

      def install
        bin.install "gogtfobins"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.14.0/gogtfobins_Linux_x86_64.tar.gz"
        sha256 "91158354d5a6138d0a39b9ca1549769f28133774817cf323a8bf7404b8f6b837"

        def install
          bin.install "gogtfobins"
        end
      end
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/juliendoutre/gogtfobins/releases/download/v1.14.0/gogtfobins_Linux_arm64.tar.gz"
        sha256 "62cab1390819cb32874bde50d8f10284060e7b2c68ffe702a33bf67f68625753"

        def install
          bin.install "gogtfobins"
        end
      end
    end
  end
end
