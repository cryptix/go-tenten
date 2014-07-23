# Converting a SOC back to a latitude and longitude:
use strict;

if ( $#ARGV != 0 ) {
    die "Usage: from-soc <10-digit-soc>";
}

my $soc = uc($ARGV[0]);
$soc =~ tr/IOSZ/1052/;

my $alphabet = 'ABCDEFGHJKLMNPQRTUVWXY0123456789';

my $soc_num = 0;

foreach my $letter (split(//, $soc)) {
    $soc_num *= 32;
    $alphabet =~ /(.*)$letter/;
    $soc_num += length($1);
}

my $p = int($soc_num / 128);
my $check = $soc_num % 128;

my $lon = $p % 3600000;
my $lat = int($p / 3600000);

$lat /= 10000;
$lon /= 10000;

$lat -= 90;
$lon -= 180;

my @primes = ( 2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37 );

my $c = 0;

foreach my $prime (@primes) {
    $c += ($p % 32) * $prime;
    $p = int($p / 32);
}

$c %= 127;

if ( $check != $c ) {
    die "Incorrect SOC";
} else {
    print "$lat $lon\n";
}
