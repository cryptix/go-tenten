# Converting a latitude and longitude to a SOC:
use strict;

if ( $#ARGV != 1 ) {
    die "Usage: to-soc  ";
}

my ( $lat, $lon ) = @ARGV;

my $alpha = 'ABCDEFGHJKLMNPQRTUVWXY0123456789';
my @alphabet = split(//,$alpha);

$lat += 90;
$lon += 180;

$lat *= 10000;
$lon *= 10000;

my $p = $lat * 3600000 + $lon;

my $soc_num = $p * 128;

my @primes = ( 2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37 );

my $c = 0;

foreach my $prime (@primes) {
    $c += ($p % 32) * $prime;
    $p = int($p / 32);
}

$c %= 127;

$soc_num += $c;

my $digits = 10;

my $soc = '';

while ( $digits > 0 ) {
    my $d = $soc_num % 32;
    $soc = $alphabet[$d] . $soc;
    $soc_num = int($soc_num/32);
    --$digits;
}

print "$soc\n";
