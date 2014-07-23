# Converting a SOC back to a latitude and longitude:
use strict;

if ( $#ARGV != 0 ) {
    die "Usage: from-soc <12-digit-soc>";
}

my $soc = uc($ARGV[0]);
$soc =~ s/ //g;

my $alpha = 'ABCDEFGHJKMNPQRVWXY0123456789';
my $base = length($alpha);

my $soc_num = 0;

foreach my $letter (split(//, $soc)) {
    $soc_num *= $base;
    $alpha =~ /(.*)$letter/;
    $soc_num += length($1);
    print $soc_num . "\n";
}

my $p = int($soc_num / $base);
my $check = $soc_num % $base;

my $lon = $p % 3600000;
my $lat = int($p / 3600000);

$lat /= 10000;
$lon /= 10000;

$lat -= 90;
$lon -= 180;


my $c = 0;

for (my $i = 1; $i < 10; $i++) {
    $c += ($p % $base) * $i;
    $p = int($p / $base);
}

$c %= $base;

if ( $check != $c ) {
    die "Incorrect SOC - got $check wanted $c";
} else {
    print "$lat $lon\n";
}
