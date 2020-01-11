use strict;
use warnings;

my @arr = ("Ali", "Mehmet", "Zeynep", "Ay�e");

my $size = @arr;

print $size, "\n";

my ($a, $b, $c, $d) = @arr;

print "$a $b $c $d\n";

my @arr2 = qw/This is just a sentence/;

print "@arr2\n";
print "$arr2[2]\n";

my @matrix = ((1,2,3),(1,2,3),(1,2,3));

print "@matrix\n";
print "$matrix[2]\n";