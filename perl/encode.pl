# Converting a latitude and longitude to a SOC:
use strict;

if ( $#ARGV != 1 ) {
    die "Usage: to-soc  ";
}

my ( $lat, $lon ) = @ARGV;

my $alpha = 'ABCDEFGHJKMNPQRVWXY0123456789';
my $base = length($alpha);
my @alphabet = split(//,$alpha);

$lat += 90;
$lon += 180;

$lat *= 10000;
$lon *= 10000;

my $p = $lat * 3600000 + $lon;

my $soc_num = $p * $base;

my $c = 0;

for (my $i = 1; $i < 10; $i++) {
    $c += ($p % $base) * $i;
    $p = int($p / $base);
}

$c %= $base;

$soc_num += $c;

my $soc = '';

for (my $i = 0; $i < 10; $i++) {
    my $d = $soc_num % $base;

    if (($i == 4) || ($i == 7)) {
      $soc = " " . $soc;
    }

    $soc = $alphabet[$d] . $soc;
    $soc_num = int($soc_num/$base);
}

print "$soc\n";
